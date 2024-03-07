using System;

class Program
{
  static void Main(string[] args)
  {

    char c = Convert.ToChar(Console.ReadLine());
    char c1 = Convert.ToChar(Console.ReadLine());
    char temp;
    if (c1 < c)
    {
      temp = c;
      c = c1;
      c1 = temp;
    }
    for (; c <= c1; c++)
    {
      System.Console.Write($"{(char)(c - 32)}{c} ");
    }
  }
}
