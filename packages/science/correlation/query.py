import argparse
import textwrap
import sys
import json
import sqlparse

def create_unique_abbreviations(names):
    abbr_dict = {}
    for name in names:
        base_abbr = ''.join(word[0].upper() for word in name.split())[:3]
        abbr = base_abbr
        counter = 1
        while abbr in abbr_dict.values():
            abbr = f"{base_abbr[:2]}{counter}"
            counter += 1
        abbr_dict[name] = abbr
    return abbr_dict

def generate_sql_query(chain_protocols, duration):
    chain_protocols = {chain.lower(): [protocol.lower() for protocol in protocols] 
                       for chain, protocols in chain_protocols.items()}

    all_chains = list(chain_protocols.keys())
    all_protocols = sorted(set(protocol for protocols in chain_protocols.values() for protocol in protocols))

    chain_abbr = create_unique_abbreviations(all_chains)
    protocol_abbr = create_unique_abbreviations(all_protocols)

    cte_statements = []
    all_columns = []

    for chain in all_chains:
        for protocol in all_protocols:
            all_columns.append(f'"{chain_abbr[chain]}{protocol_abbr[protocol]}"')

    for chain, protocols in chain_protocols.items():
        case_statements = []
        for other_chain in all_chains:
            for protocol in all_protocols:
                if other_chain == chain and protocol in protocols:
                    case_statements.append(f"""
    COALESCE(SUM(CASE WHEN lower(contract_name) LIKE '%{protocol}%' THEN 1 ELSE 0 END), 0) AS "{chain_abbr[chain]}{protocol_abbr[protocol]}" """)
                else:
                    case_statements.append(f"""0 AS "{chain_abbr[other_chain]}{protocol_abbr[protocol]}" """)

        case_statements_str = ',\n'.join(case_statements)
        like_any_clause = ', '.join(f"'%{protocol}%'" for protocol in protocols)

        cte_statements.append(f"""
{chain}_behavior AS (
  SELECT
    origin_from_address AS user_,
    {case_statements_str}
  FROM
    {chain}.core.ez_decoded_event_logs
  WHERE
    block_timestamp >= DATEADD(day, -{duration}, CURRENT_DATE)
    AND contract_name IS NOT NULL
    AND lower(contract_name) LIKE ANY ({like_any_clause})
  GROUP BY
    user_
)""")

    cte_statements_str = ',\n'.join(cte_statements)
    all_columns_str = ',\n    '.join(all_columns)
    
    combined_behavior_ctes = []
    for chain in all_chains:
        combined_behavior_ctes.append(f"""
  SELECT
    user_,
    {all_columns_str}
  FROM
    {chain}_behavior""")

    combined_behavior_str = '\nUNION ALL\n'.join(combined_behavior_ctes)

    sum_statements = [f'SUM({col}) AS {col}' for col in all_columns]
    sum_statements_str = ',\n  '.join(sum_statements)

    base_query = f"""
WITH {cte_statements_str},
combined_behavior AS ({combined_behavior_str}
)
SELECT
  user_,
  {sum_statements_str}
FROM
  combined_behavior
GROUP BY
  user_
"""

    final_query = f"""
WITH base_query AS (
{base_query}
),
source_table AS (
  SELECT 
    {',\n  '.join([f'{col}' for col in all_columns])},
    ROW_NUMBER() OVER (ORDER BY user_) AS "N"
  FROM base_query
),
group_rows AS (
  SELECT 
    CEIL("N" / 1000) AS group_num,
    ARRAY_AGG(CONCAT({", ',', ".join([f'{col}' for col in all_columns])})) AS data_array
  FROM source_table
  GROUP BY group_num
)
SELECT 
  group_num,
  data_array
FROM group_rows
ORDER BY group_num;
"""

    return final_query

def main():
    parser = argparse.ArgumentParser(description='Generate SQL query for multi-chain protocol analysis.')
    parser.add_argument('--chains', nargs='+', required=True, help='List of chains to analyze')
    parser.add_argument('--protocols', nargs='+', action='append', required=True, 
                        help='List of protocols for each chain. Use multiple times, once for each chain.')
    parser.add_argument('--duration', default='30', help='Duration of the query in days (default: 30)')
    parser.add_argument('--output', default="correlation/query.txt", help='Output file name (default: correlation/query.txt)')

    args = parser.parse_args()

    if len(args.chains) != len(args.protocols):
        print("Error: The number of --chains arguments must match the number of --protocols arguments.")
        sys.exit(1)

    chain_protocols = dict(zip(args.chains, args.protocols))

    query = generate_sql_query(chain_protocols, args.duration)
    
    formatted_query = sqlparse.format(query, reindent=True, keyword_case='upper')

    with open(args.output, 'w') as f:
        f.write(formatted_query)
    print(f"Formatted query has been written to {args.output}")

if __name__ == '__main__':
    main()