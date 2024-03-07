using System;

class Program
{
  static void Main()
  {
    int y = Int32.Parse(Console.ReadLine());
    int n = Int32.Parse(Console.ReadLine());
    int counter = 0;
    while (y >= 0)
    {
      y -= n + n / 3;
      counter++;
    }
    Console.WriteLine(counter);
  }
}
