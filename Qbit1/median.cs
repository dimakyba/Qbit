using System;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] data = Console.ReadLine().Trim().Split();
            int a = int.Parse(data[0]);
            int b = int.Parse(data[1]);
            int c = int.Parse(data[2]);
            int p = (a + b + c) / 2;
            if (a + b > c && a + c > b && b + c > a)
            {
                Console.WriteLine(0.5 * Math.Sqrt(2 * b * b + 2 * c * c - a * a));
                Console.WriteLine(0.5 * Math.Sqrt(2 * a * a + 2 * c * c - b * b));
                Console.WriteLine(0.5 * Math.Sqrt(2 * a * a + 2 * b * b - c * c));
            }
        }
    }
}
