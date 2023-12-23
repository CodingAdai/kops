import {request} from "@/utils/service";

export interface UpdateKubeconfigContentParam {
  content: string
}

export type LoginResponseData = ApiResponseData<boolean>

export function updateKubeconfigContent(data: UpdateKubeconfigContentParam) {
  return request<LoginResponseData>({
    baseURL: "http://localhost:9091",
    url: "/api/k8s/kubeconfig/content",
    method: "post",
    data
  })
}
