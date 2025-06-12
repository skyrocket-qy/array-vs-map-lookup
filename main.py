import pandas as pd
import matplotlib.pyplot as plt

# Load data
df = pd.read_csv("benchmark.csv")

# Plot
plt.figure(figsize=(10, 6))
plt.plot(df["n"], df["array_lookup"], label="Array Lookup (O(n))", color="blue")
plt.plot(df["n"], df["map_lookup"], label="Map Lookup (O(1))", color="green")
plt.xlabel("Number of Elements (n)")
plt.ylabel("Average Lookup Time (seconds)")
plt.title("Go Benchmark: Array vs. Map Lookup Time")
plt.legend()
plt.grid(True)
plt.tight_layout()
plt.savefig("go_array_vs_map_lookup.png")
plt.show()
