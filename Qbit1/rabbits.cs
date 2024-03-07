using System;

class Program
{
  static void Main()
  {
    string[] data = Console.ReadLine().Trim().Split();
    int n = Int32.Parse(data[0]);
    int m = Int32.Parse(data[1]);

    Console.WriteLine(Math.Ceiling((double)m / (double)n));
  }
}
