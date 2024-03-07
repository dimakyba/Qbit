using System;

class Program
{
  static void Main()
  {
    string[] data = Console.ReadLine().Trim().Split();
    int a = Int32.Parse(data[0]);
    int b = Int32.Parse(data[1]);
    int n = Int32.Parse(data[2]);
    float cash = (float)(n * (a + b / 100.0));
    Console.WriteLine($"{Math.Floor(cash)} {(int)Math.Round((cash - Math.Floor(cash)) * 100)}");

  }
}
