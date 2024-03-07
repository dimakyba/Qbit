using System;

class Program
{
  static void Main()
  {
    int h = int.Parse(Console.ReadLine());
    int a = int.Parse(Console.ReadLine());
    int b = int.Parse(Console.ReadLine());
    int upPerDay = a - b;
    int days = (int)Math.Ceiling((double)(h - a) / upPerDay) + 1;

    Console.WriteLine(days);
  }
}
