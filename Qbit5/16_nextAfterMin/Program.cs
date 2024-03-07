using System;
using System.Linq;

class Program
{
  static void Main()
  {
    int n = int.Parse(Console.ReadLine());
    int[] array = Console.ReadLine().Split(' ').Select(int.Parse).ToArray();

    FindNextAfterMin(array);
  }

  static void FindNextAfterMin(int[] array)
  {
    int min = int.MaxValue;
    int minIndex = -1;

    for (int i = 0; i < array.Length; i++)
    {
      if (array[i] < min)
      {
        min = array[i];
        minIndex = i;
      }
    }

    int nextAfterMin = int.MaxValue;
    int nextAfterMinIndex = -1;

    for (int i = 0; i < array.Length; i++)
    {
      if (array[i] > min && array[i] < nextAfterMin)
      {
        nextAfterMin = array[i];
        nextAfterMinIndex = i;
      }
    }

    Console.WriteLine($"{nextAfterMin} {nextAfterMinIndex + 1}");
  }
}
