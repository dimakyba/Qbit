using System;

class Program
{
  static void Main()
  {
    uint s = UInt32.Parse(Console.ReadLine());
    uint hours = s / 3600;
    s %= 3600;
    uint minutes = s / 60;
    uint seconds = s % 60;

    if (hours >= 24)
    {
      hours %= 24;
    }

    Console.WriteLine($"{hours:d2}:{minutes:d2}:{seconds:d2}");
  }
}
