import pandas as pd
import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt

df = pd.read_csv("correlation/actions.csv")

y = df.iloc[:, 2:10]
y2 = df.iloc[:, 4:10]

def sort_correlation_matrix(data):
    corr = data.corr()
    corr_array = corr.abs().to_numpy()
    np.fill_diagonal(corr_array, 0)
    indices = np.argsort(corr_array.sum(axis=1))[::-1]
    corr_sorted = corr.iloc[indices, indices]
    return corr_sorted

z = sort_correlation_matrix(y)
z2 = sort_correlation_matrix(y2)

def plot_correlation(corr_matrix, title):
    plt.figure(figsize=(12, 10))
    mask = np.triu(np.ones_like(corr_matrix, dtype=bool)) | (np.abs(corr_matrix) < 1e-10)
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

plot_correlation(z, "Sorted Correlation Matrix (Columns 3-10)")
plot_correlation(z2, "Sorted Correlation Matrix (Columns 5-10)")