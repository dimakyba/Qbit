using System;
using System.Linq;

class Program
{
  static void Main()
  {
    int[] array = Console.ReadLine().Split(' ').Select(int.Parse).ToArray();
    string[] indexes = Console.ReadLine().Trim().Split();
    int f = int.Parse(indexes[0]);
    int k = int.Parse(indexes[1]);

    int adjustedIndex = k - f;
    if (adjustedIndex >= 0 && adjustedIndex < array.Length)
    {
      Console.WriteLine(array[adjustedIndex]);
    }
    else
    {
      Console.WriteLine("Error");
    }
  }
}
