// variables.js
export const httpDev = 'http://localhost:8081'
export const httpProd = 'http://localhost:8081'
export const httpBaseURL = process.env.NODE_ENV === 'production' ? 'http://localhost:8081' : 'http://localhost:8081'
export const settings = {
  some: 'Settings'
}