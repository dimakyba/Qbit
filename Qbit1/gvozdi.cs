using System;

class Program
{
  static void Main()
  {
    string[] data = Console.ReadLine().Trim().Split();
    int a = Int32.Parse(data[0]);
    int b = Int32.Parse(data[1]);
    int res = 37 * a + b;
    Console.WriteLine($"{res / 23} {res % 23}");
  }
}
