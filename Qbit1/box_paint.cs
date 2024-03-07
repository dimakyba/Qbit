using System;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] data = Console.ReadLine().Trim().Split();
            float a = Convert.ToSingle(data[0]);
            float b = Convert.ToSingle(data[1]);
            float c = Convert.ToSingle(data[2]);
            float x = Convert.ToSingle(data[3]);
            float area = 2 * (a * b + b * c + a * c);
            Console.WriteLine(Math.Ceiling(x * area /10));



        }
    }
}
