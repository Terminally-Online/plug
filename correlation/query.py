import argparse
import textwrap
import sys

def generate_sql_query(protocols, start_date):
    # Create the CASE statements for each protocol
    case_statements = []
    sum_statements = []
    for protocol in protocols:
        case_statements.append(f"""
    case
      when contract_name ILIKE ('%{protocol}%') then 1
      else 0
    end as {protocol}_flag""")
        sum_statements.append(f"sum({protocol}_flag) as n_{protocol}")

    # Join the statements
    case_statements_str = ',\n'.join(case_statements)
    sum_statements_str = ',\n  '.join(sum_statements)

    # Generate the ILIKE ANY clause
    ilike_any_clause = ', '.join(f"'%{protocol}%'" for protocol in protocols)

    # Create the full SQL query
    sql_query = f"""
    with user_behavior AS (
      select
        origin_from_address as user_,
        date_trunc('month', block_timestamp) as month_,
        tx_hash,
    {case_statements_str}
      from
        ethereum.core.ez_decoded_event_logs
      where
        block_timestamp >= '{start_date}'
        and contract_name IS NOT NULL
        and contract_name ILIKE ANY (
          {ilike_any_clause}
        )
    )
    select
      user_,
      month_,
      count(distinct tx_hash) as n_tx,
      {sum_statements_str}
    from
      user_behavior
    group by
      user_,
      month_
    """

    return textwrap.dedent(sql_query).strip()

def main():
    parser = argparse.ArgumentParser(description='Generate SQL query for protocol analysis.')
    parser.add_argument('protocols', nargs='+', help='List of protocols to include in the query')
    parser.add_argument('--start-date', default='2024-05-01', help='Start date for the query (default: 2024-05-01)')
    parser.add_argument('--output', default="correlation/query.txt", help='Output file name (optional)')

    args = parser.parse_args()

    query = generate_sql_query(args.protocols, args.start_date)

    if args.output:
        with open(args.output, 'w') as f:
            f.write(query)
        print(f"Query has been written to {args.output}")
    else:
        print(query)

if __name__ == '__main__':
    main()