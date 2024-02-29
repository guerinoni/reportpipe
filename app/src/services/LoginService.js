async function login(email, password) {
    const body = {
        email,
        password
    }
    return await fetch('https://api.example.com/...', body)
}

async function register(email, password) {
    const body = {
        email,
        password
    }
    return await fetch('https://api.example.com/...', body)
}

const LoginService = {
    login,
    register
}

export default LoginService