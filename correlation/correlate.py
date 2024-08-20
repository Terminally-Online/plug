import pandas as pd
import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt

def sort_correlation_matrix(data):
    corr = data.corr()
    corr_array = corr.abs().to_numpy()
    np.fill_diagonal(corr_array, 0)
    indices = np.argsort(corr_array.sum(axis=1))[::-1]
    corr_sorted = corr.iloc[indices, indices]
    return corr_sorted

def plot_correlation(corr_matrix, title):
    plt.figure(figsize=(max(12, len(corr_matrix) * 0.8), max(10, len(corr_matrix) * 0.7)))
    
    mask = np.triu(np.ones_like(corr_matrix, dtype=bool))
    
    sns.heatmap(corr_matrix, 
                mask=mask,
                annot=True, 
                cmap=sns.diverging_palette(220, 10, as_cmap=True),
                vmin=0, 
                vmax=1, 
                square=True, 
                cbar_kws={"shrink": .8},
                fmt=".2f")
    
    plt.title(title)
    plt.xticks(rotation=45, ha='right')
    plt.yticks(rotation=0)
    plt.tight_layout()
    plt.show()

df = pd.read_csv("correlation/actions.csv")

analysis_columns = df.columns[2:]

non_zero_columns = [col for col in analysis_columns if df[col].sum() != 0]
excluded_columns = set(analysis_columns) - set(non_zero_columns)

if excluded_columns:
    print("Warning: The following columns were excluded due to having all zero values:")
    for col in excluded_columns:
        print(f"- {col}")

z = sort_correlation_matrix(df[non_zero_columns])
plot_correlation(z, f"Protocol and Transaction Correlation Matrix ({len(non_zero_columns)} columns)")