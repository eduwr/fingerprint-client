import { UAParser, IResult } from "ua-parser-js"


export type FingerprintData = {
  screen: string
  timezone: string
  language: string
  maxTouchPoints: number
  gpu: string
  canvas: string
}

export const getFingerprint = async () => {
  const uaResult = UAParser(window.navigator.userAgent)
  const data: FingerprintData = {
    screen: `${window.screen.width}x${window.screen.height}`,
    timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
    language: navigator.language,
    maxTouchPoints: navigator.maxTouchPoints || 0,
    gpu: getWebGLFingerprint(),
    canvas: getCanvasFingerprint(),
  }

  console.log({ data })

  const response = await fetch(`${import.meta.env.VITE_API_URL}/fingerprint`, {
    body: JSON.stringify(data),
    method: "post",
    headers: {
      "Content-Type": "application/json"
    }
  })

  await response.json()
  console.log(response)

  return data
}

const getWebGLFingerprint = () => {
  const canvas = document.createElement("canvas")
  const gl = (canvas.getContext("webgl") ||
    canvas.getContext("experimental-webgl")) as WebGLRenderingContext | null

  if (!gl) return "unsupported"

  const debugInfo = gl.getExtension("WEBGL_debug_renderer_info")
  const vendor = debugInfo
    ? gl.getParameter(debugInfo.UNMASKED_VENDOR_WEBGL)
    : "Unknown"
  const renderer = debugInfo
    ? gl.getParameter(debugInfo.UNMASKED_RENDERER_WEBGL)
    : "Unknown"

  return `${vendor} - ${renderer}`
}

const getCanvasFingerprint = (): string => {
  const canvas = document.createElement("canvas")
  const ctx = canvas.getContext("2d")

  if (!ctx) return "unsupported"

  ctx.textBaseline = "top"
  ctx.font = "16px Arial"
  ctx.fillText("Fingerprinting Test!", 10, 10)

  return canvas.toDataURL()
}
