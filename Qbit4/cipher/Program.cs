using System;
using System.Text;

class Program
{
  static void Main(string[] args)
  {
    int n = int.Parse(Console.ReadLine());
    string s = Console.ReadLine();

    foreach (char c in s)
    {
      System.Console.Write($"{(int)c} ");
    }
  }
}
