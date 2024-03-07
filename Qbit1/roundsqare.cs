using System;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            int radius = Int32.Parse(Console.ReadLine());
            Console.WriteLine(Math.PI* radius* radius);
            Console.WriteLine(2 * Math.PI * radius);
        }
    }
}
