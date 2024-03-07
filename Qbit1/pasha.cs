using System;
using System.Linq;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            const int sumOfFive = 15;
            int[] numbers = Console.ReadLine().Split(' ').Select(int.Parse).ToArray();
            Console.WriteLine(sumOfFive - numbers.Sum());
        }
    }
}
