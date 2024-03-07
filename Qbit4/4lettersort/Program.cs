using System;
using System.Linq;

class Program
{
  static void Main(string[] args)
  {
    char[] letters = Console.ReadLine().Split(' ').Select(char.Parse).ToArray();
    Array.Sort(letters);
    foreach (char letter in letters)
    {
      Console.Write($"{letter} ");
    }
  }
}
