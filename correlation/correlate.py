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
    
    mask = np.abs(corr_matrix) < 0.0
    
    sns.heatmap(corr_matrix, 
                mask=mask,
                annot=True, 
                cmap=sns.diverging_palette(220, 10, as_cmap=True),
                vmin=-1, 
                vmax=1, 
                square=True, 
                cbar_kws={"shrink": .8},
                fmt=".2f")
    
    plt.title(title)
    plt.xticks(rotation=45, ha='right')
    plt.yticks(rotation=0)
    plt.tight_layout()
    plt.show()

# Read the CSV file
df = pd.read_csv("correlation/actions.csv")

# Select all columns starting from the third column (index 2)
analysis_columns = df.columns[2:]

# Generate and plot the correlation matrix
z = sort_correlation_matrix(df[analysis_columns])
plot_correlation(z, f"Protocol and Transaction Correlation Matrix ({len(analysis_columns)} columns)")