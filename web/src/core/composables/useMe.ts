import { jwtDecode } from 'jwt-decode'

export function useMe() {
  function getDecodedToken() {
    const token = localStorage.getItem('token')
    if (!token) {
      return null
    }
    try {
      return jwtDecode(token) as any
    } catch (error) {
      console.error('Invalid token:', error)
      return null
    }
  }

  return {
    getDecodedToken,
  }
}
