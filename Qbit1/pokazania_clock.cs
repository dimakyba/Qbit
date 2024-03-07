using System;

class Program
{
  static void Main()
  {
    string[] data = Console.ReadLine().Trim().Split();
    int h = Int32.Parse(data[0]);
    int m = Int32.Parse(data[1]);
    int s = Int32.Parse(data[2]);
    string[] data2 = Console.ReadLine().Trim().Split();
    int h1 = Int32.Parse(data2[0]);
    int m1 = Int32.Parse(data2[1]);
    int s1 = Int32.Parse(data2[2]);

    int totalSecs = s + s1 + m * 60 + m1 * 60 + h * 3600 + h1 * 3600;

    totalSecs %= 86400;
    int hours = totalSecs / 3600;
    totalSecs %= 3600;

    int mins = totalSecs / 60;
    int secs = totalSecs % 60;

    Console.WriteLine($"{hours} {mins} {secs}");
  }
}
