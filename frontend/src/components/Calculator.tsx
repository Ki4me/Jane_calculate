'use client';

import { useState } from 'react';
import { calculatorClient } from '../api/calculator';

export default function Calculator() {
  const [a, setA] = useState<number>(0);
  const [b, setB] = useState<number>(0);
  const [operation, setOperation] = useState<string>('add');
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleCalculate = async () => {
    try {
      setError(null);
      const response = await calculatorClient.calculate({
        a,
        b,
        operation,
      });
      setResult(response.result);
    } catch (err: any) {
      setError(err.message || '计算出错');
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 py-6 flex flex-col justify-center sm:py-12">
      <div className="relative py-3 sm:max-w-xl sm:mx-auto">
        <div className="relative px-4 py-10 bg-white shadow-lg sm:rounded-3xl sm:p-20">
          <div className="max-w-md mx-auto">
            <div className="divide-y divide-gray-200">
              <div className="py-8 text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7">
                <div className="flex flex-col space-y-4">
                  <input
                    type="number"
                    value={a}
                    onChange={(e) => setA(Number(e.target.value))}
                    className="px-4 py-2 border rounded-md"
                    placeholder="第一个数字"
                  />
                  <select
                    value={operation}
                    onChange={(e) => setOperation(e.target.value)}
                    className="px-4 py-2 border rounded-md"
                  >
                    <option value="add">加法</option>
                    <option value="subtract">减法</option>
                    <option value="multiply">乘法</option>
                    <option value="divide">除法</option>
                  </select>
                  <input
                    type="number"
                    value={b}
                    onChange={(e) => setB(Number(e.target.value))}
                    className="px-4 py-2 border rounded-md"
                    placeholder="第二个数字"
                  />
                  <button
                    onClick={handleCalculate}
                    className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
                  >
                    计算
                  </button>
                  {result !== null && (
                    <div className="mt-4 p-4 bg-green-100 rounded-md">
                      结果: {result}
                    </div>
                  )}
                  {error && (
                    <div className="mt-4 p-4 bg-red-100 rounded-md text-red-700">
                      错误: {error}
                    </div>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
} 