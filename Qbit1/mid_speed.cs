using System;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] data = Console.ReadLine().Trim().Split();
            double v1 = Convert.ToDouble(data[0]);
            double v2 = Convert.ToDouble(data[1]);
            double v3 = Convert.ToDouble(data[2]);
            double v4 = Convert.ToDouble(data[3]);

            Console.WriteLine((v1 + v2) / 2);
            Console.WriteLine((2 * v3 * v4) / (v3 + v4));



        }
    }
}
