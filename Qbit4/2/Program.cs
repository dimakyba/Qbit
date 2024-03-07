using System;

class Program
{
  static void Main()
  {
    string s = Console.ReadLine();
    string[] data = Console.ReadLine().Split(' ');

    int n = int.Parse(data[0]);
    int m = int.Parse(data[1]);

    n--; m--;

    if (n < 0)
      n = 0;

    if (m >= s.Length)
      m = s.Length - 1;

    string resultString = s.Remove(n, m - n + 1);

    Console.WriteLine(resultString);
  }

}
