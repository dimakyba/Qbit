using System;

class Program
{
  static void Main()
  {
    int x = 0, y = 0;
    int dx = 0, dy = 1;
    string c = Console.ReadLine();

    while (!string.IsNullOrEmpty(c))
    {
      if (c == "R")
      {
        int temp = dx;
        dx = dy;
        dy = -temp;
      }
      else if (c == "L")
      {
        int temp = dx;
        dx = -dy;
        dy = temp;
      }
      else
      {
        x += dx;
        y += dy;
      }
      c = Console.ReadLine();
    }

    Console.WriteLine($"{x} {y}");
  }
}
