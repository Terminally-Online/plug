import argparse
import textwrap
import sys
import json
import sqlparse

def generate_sql_query(chain_protocols, start_date):
    # Standardize all chain and protocol names to lowercase
    chain_protocols = {chain.lower(): [protocol.lower() for protocol in protocols] 
                       for chain, protocols in chain_protocols.items()}

    all_chains = list(chain_protocols.keys())
    all_protocols = sorted(set(protocol for protocols in chain_protocols.values() for protocol in protocols))

    cte_statements = []
    all_flag_columns = []

    for chain in all_chains:
        all_flag_columns.append(f'"{chain}_tx"')
        for protocol in all_protocols:
            all_flag_columns.append(f'"{chain}_{protocol}_flag"')

    for chain, protocols in chain_protocols.items():
        case_statements = []
        for other_chain in all_chains:
            if other_chain == chain:
                case_statements.append(f'COUNT(DISTINCT tx_hash) AS "{chain}_tx"')
            else:
                case_statements.append(f'0 AS "{other_chain}_tx"')
            for protocol in all_protocols:
                if other_chain == chain and protocol in protocols:
                    case_statements.append(f"""
    COALESCE(SUM(CASE WHEN lower(contract_name) LIKE '%{protocol}%' THEN 1 ELSE 0 END), 0) AS "{other_chain}_{protocol}_flag" """)
                else:
                    case_statements.append(f"""0 AS "{other_chain}_{protocol}_flag" """)

        case_statements_str = ',\n'.join(case_statements)
        like_any_clause = ', '.join(f"'%{protocol}%'" for protocol in protocols)

        cte_statements.append(f"""
{chain}_behavior AS (
  SELECT
    origin_from_address AS user_,
    date_trunc('month', block_timestamp) AS month_,
    '{chain}' AS chain,
    {case_statements_str}
  FROM
    {chain}.core.ez_decoded_event_logs
  WHERE
    block_timestamp >= '{start_date}'
    AND contract_name IS NOT NULL
    AND lower(contract_name) LIKE ANY ({like_any_clause})
  GROUP BY
    user_, month_
)""")

    cte_statements_str = ',\n'.join(cte_statements)
    all_flag_columns_str = ',\n    '.join(all_flag_columns)
    
    combined_behavior_ctes = []
    for chain in all_chains:
        combined_behavior_ctes.append(f"""
  SELECT
    user_,
    month_,
    {all_flag_columns_str}
  FROM
    {chain}_behavior""")

    combined_behavior_str = '\nUNION ALL\n'.join(combined_behavior_ctes)

    sum_statements = [f'COALESCE(SUM({col}), 0) AS {col}' for col in all_flag_columns]
    sum_statements_str = ',\n  '.join(sum_statements)

    sql_query = f"""
WITH {cte_statements_str},
combined_behavior AS ({combined_behavior_str}
)
SELECT
  user_,
  month_,
  {sum_statements_str}
FROM
  combined_behavior
GROUP BY
  user_,
  month_
"""

    return sql_query

def main():
    parser = argparse.ArgumentParser(description='Generate SQL query for multi-chain protocol analysis.')
    parser.add_argument('--chains', nargs='+', required=True, help='List of chains to analyze')
    parser.add_argument('--protocols', nargs='+', action='append', required=True, 
                        help='List of protocols for each chain. Use multiple times, once for each chain.')
    parser.add_argument('--start-date', default='2024-05-01', help='Start date for the query (default: 2024-05-01)')
    parser.add_argument('--output', default="correlation/query.txt", help='Output file name (default: correlation/query.txt)')

    args = parser.parse_args()

    if len(args.chains) != len(args.protocols):
        print("Error: The number of --chains arguments must match the number of --protocols arguments.")
        sys.exit(1)

    chain_protocols = dict(zip(args.chains, args.protocols))

    query = generate_sql_query(chain_protocols, args.start_date)
    
    formatted_query = sqlparse.format(query, reindent=True, keyword_case='upper')

    with open(args.output, 'w') as f:
        f.write(formatted_query)
    print(f"Formatted query has been written to {args.output}")

if __name__ == '__main__':
    main()