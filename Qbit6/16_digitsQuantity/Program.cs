using System;

class Program
{
  static void Main()
  {
    string[] inputNumbers = Console.ReadLine().Split();
    int[] counts = new int[9];
    foreach (string number in inputNumbers)
    {
      if (number == "0")
        break;

      int digit = int.Parse(number);
      counts[digit - 1]++;
    }
    Console.WriteLine(string.Join(" ", counts));
  }
}
