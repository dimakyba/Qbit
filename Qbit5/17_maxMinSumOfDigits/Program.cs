using System;
using System.Linq;

class Program
{

  static void Main()
  {
    int N = int.Parse(Console.ReadLine());
    int[] array = Console.ReadLine().Split().Select(int.Parse).ToArray();

    int[] result = FindMinMaxSumDigits(array);
    Console.WriteLine(result[0]);
    Console.WriteLine(result[1]);
  }
  static int SumOfDigits(int num)
  {
    return Math.Abs(num).ToString().Sum(c => int.Parse(c.ToString()));
  }

  static int[] FindMinMaxSumDigits(int[] arr)
  {
    int minSumDigits = int.MaxValue;
    int maxSumDigits = int.MinValue;

    int[] result = new int[2];

    foreach (int num in arr)
    {
      int currentSumDigits = SumOfDigits(num);

      if (currentSumDigits < minSumDigits || (currentSumDigits == minSumDigits && num < result[0]))
      {
        minSumDigits = currentSumDigits;
        result[0] = num;
      }

      if (currentSumDigits > maxSumDigits || (currentSumDigits == maxSumDigits && num > result[1]))
      {
        maxSumDigits = currentSumDigits;
        result[1] = num;
      }
    }

    return result;
  }



}
