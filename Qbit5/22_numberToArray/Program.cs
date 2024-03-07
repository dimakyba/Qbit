using System;
using System.Collections.Generic;

class Program
{
  static void Main(string[] args)
  {
    int k = int.Parse(Console.ReadLine());

    int[] arr = ExtractDigits(k);

    PrintArray(arr);
  }

  static int[] ExtractDigits(int n)
  {
    if (n == 0)
    {
      return 0;
    }

    var digits = new List<int>();

    while (n != 0)
    {
      int digit = n % 10;
      digits.Insert(0, digit);
      n /= 10;
    }

    return digits.ToArray();
  }

  static void PrintArray(int[] arr)
  {
    foreach (int v in arr)
    {
      Console.Write($"{v} ");
    }
    Console.WriteLine();
  }
}
