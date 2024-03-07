using System;

class Program
{
  static void Main()
  {
    uint n = UInt32.Parse(Console.ReadLine());
    if ((n + 5) % 7 == 0)
    {
      Console.WriteLine(7);
    }
    else
    {
      Console.WriteLine((n + 5) % 7);
    }
  }
}
