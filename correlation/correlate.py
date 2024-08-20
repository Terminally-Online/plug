import pandas as pd
import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt
import matplotlib.colors as mcolors

def sort_correlation_matrix(data):
    corr = data.corr()
    
    np.fill_diagonal(corr.values, 0)
    column_sums = corr.sum()
    sorted_columns = column_sums.sort_values(ascending=False).index
    
    corr_sorted = corr.loc[sorted_columns, sorted_columns]
    
    return corr_sorted

def plot_correlation(corr_matrix, title):
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
                annot_kws={"size": 6})
    
    ax.set_title(title, fontsize=12, pad=20)
    ax.set_xticklabels(ax.get_xticklabels(), rotation=90, ha='right', fontsize=8)
    ax.set_yticklabels(ax.get_yticklabels(), rotation=0, fontsize=8)
    
    # Adjust layout to prevent cutting off labels
    plt.tight_layout()
    
    # Add padding
    fig.subplots_adjust(left=0.2, right=0.9, top=0.9, bottom=0.15)
    
    # Adjust x-axis labels
    ax.tick_params(axis='x', which='major', pad=5)
    
    plt.show()

# Read the CSV file
df = pd.read_csv("correlation/actions.csv")

# Select all columns starting from the third column (index 2)
analysis_columns = df.columns[2:]

# Exclude columns with all zero values
non_zero_columns = [col for col in analysis_columns if df[col].sum() != 0]
excluded_columns = set(analysis_columns) - set(non_zero_columns)

if excluded_columns:
    print("Warning: The following columns were excluded due to having all zero values:")
    for col in excluded_columns:
        print(f"- {col}")

# Generate and plot the correlation matrix
z = sort_correlation_matrix(df[non_zero_columns])
plot_correlation(z, f"Protocol and Transaction Correlation Matrix ({len(non_zero_columns)} columns)")