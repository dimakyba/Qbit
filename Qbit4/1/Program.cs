using System;
using System.Text;

class Program
{
  static void Main(string[] args)
  {
    char c = Convert.ToChar(Console.ReadLine());
    System.Console.WriteLine($"The next char for '{c}' ({(int)c}) is '{(char)(c + 1)}' ({(int)c + 1}).");
    System.Console.WriteLine($"The previous char for '{c}' ({(int)c}) is '{(char)(c - 1)}' ({(int)c - 1}).");
  }
}
