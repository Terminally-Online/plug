import pandas as pd
import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt
import json
import argparse
import sys

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

def parse_csv_data(file_path):
    df = pd.read_csv(file_path)
    
    def parse_string(s):
        return [int(x) for x in s.strip('"').split(',')]
    
    all_data = []
    for _, row in df.iterrows():
        data_array = json.loads(row['DATA_ARRAY'])
        for item in data_array:
            all_data.append(parse_string(item))
    
    return pd.DataFrame(all_data)

def remove_zero_columns(df):
    non_zero_columns = df.columns[df.sum() != 0]
    return df[non_zero_columns]

def sort_correlation_matrix(data):
    corr = data.corr()
    corr_array = corr.abs().to_numpy()
    np.fill_diagonal(corr_array, 0)
    column_sums = corr.sum()
    sorted_columns = column_sums.sort_values(ascending=False).index
    return corr.loc[sorted_columns, sorted_columns]

def plot_correlation(corr_matrix, title, labels, full_names, output_file):
    fig, ax = plt.subplots(figsize=(max(12, len(corr_matrix) * 0.8), max(10, len(corr_matrix) * 0.7)))
    
    mask = np.triu(np.ones_like(corr_matrix, dtype=bool))
    
    sns.heatmap(corr_matrix, 
                ax=ax,
                mask=mask,
                annot=True, 
                cmap=sns.diverging_palette(220, 10, as_cmap=True),
                vmin=0,
                vmax=1, 
                square=True, 
                cbar_kws={"shrink": .8},
                fmt=".2f",
                xticklabels=[full_names[label] for label in labels],
                yticklabels=[full_names[label] for label in labels])
    
    ax.set_title(title, fontsize=12, pad=20)
    ax.set_xticklabels(ax.get_xticklabels(), rotation=90, ha='right', fontsize=8)
    ax.set_yticklabels(ax.get_yticklabels(), rotation=0, fontsize=8)
    
    plt.tight_layout()
    
    plt.savefig(output_file, dpi=300, bbox_inches='tight')
    print(f"Correlation matrix image saved as {output_file}")
    
    plt.show()

def main():
    parser = argparse.ArgumentParser(description='Generate correlation matrix for multi-chain protocol analysis.')
    parser.add_argument('--chains', nargs='+', required=True, help='List of chains analyzed')
    parser.add_argument('--protocols', nargs='+', action='append', required=True, 
                        help='List of protocols for each chain. Use multiple times, once for each chain.')
    parser.add_argument('--input', default="correlation/actions.csv", help='Input CSV file (default: correlation/actions.csv)')
    parser.add_argument('--output', default="correlation/matrix.png", help='Output image file (default: correlation_matrix.png)')

    args = parser.parse_args()

    if len(args.chains) != len(args.protocols):
        print("Error: The number of --chains arguments must match the number of --protocols arguments.")
        sys.exit(1)

    chain_protocols = dict(zip(args.chains, args.protocols))
    
    chain_abbr = create_unique_abbreviations(chain_protocols.keys())
    all_protocols = sorted(set(protocol for protocols in chain_protocols.values() for protocol in protocols))
    protocol_abbr = create_unique_abbreviations(all_protocols)

    labels = []
    full_names = {}
    for chain in chain_protocols.keys():
        for protocol in all_protocols:
            abbr = f"{chain_abbr[chain]}{protocol_abbr[protocol]}"
            full_name = f"{chain.capitalize()} - {protocol.capitalize()}"
            labels.append(abbr)
            full_names[abbr] = full_name

    df = parse_csv_data(args.input)
    df_non_zero = remove_zero_columns(df)
    non_zero_indices = [i for i, col in enumerate(df.columns) if col in df_non_zero.columns]
    labels = [labels[i] for i in non_zero_indices]
    removed_columns = set(df.columns) - set(df_non_zero.columns)
    z = sort_correlation_matrix(df_non_zero)
    plot_correlation(z, f"Protocol Correlation Matrix ({len(df_non_zero.columns)} active protocols)", labels, full_names, args.output)

if __name__ == '__main__':
    main()