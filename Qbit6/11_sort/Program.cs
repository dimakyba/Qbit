using System;
using System.Linq;

namespace Program
{
  class Program
  {
    static int Sort(int[] arr, int n)
    {
      int counter = 0;
      for (int i = 0; i < n - 1; i++)
      {
        for (int j = 0; j < n - 1; j++)
        {
          if (arr[j] > arr[j + 1])
          {
            (arr[j], arr[j + 1]) = (arr[j + 1], arr[j]);
            counter++;
          }
        }
      }

      return counter;
    }

    static void Main(string[] args)
    {
      int length = int.Parse(Console.ReadLine());
      int[] array = Console.ReadLine().Split().Select(int.Parse).ToArray();
      Console.WriteLine(Sort(array, length));
    }
  }
}
