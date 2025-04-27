import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { CalculatorService } from "./api/calculator_pb";

// 创建 Connect-RPC 传输层
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  useBinaryFormat: false,
});

// 创建 Connect-RPC 客户端
export const calculatorClient = createClient(CalculatorService, transport); 