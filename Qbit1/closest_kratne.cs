using System;

class Program
{
  static void Main()
  {
    string[] data = Console.ReadLine().Trim().Split();
    int n = Int32.Parse(data[0]);
    int p = Int32.Parse(data[1]);
    int temp = n % p;
    int m;

    if (temp == 0)
    {
      m = n;
    }
    else if (n >= 0)
    {
      m = n + (p - temp);
    }
    else
    {
      m = n - temp;
    }

    Console.WriteLine(m);
  }
}
