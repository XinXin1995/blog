import request from '@/utils/request'

export function Login({username, password, captcha, captchaId}) {
    return request.post('/api/user/login', {username, password, captcha, captchaId})
}

export function Register({username, password, email}) {
    return request.post('/api/user/register', {username, password, email})
}
export function GetCaptcha() {
    return request.get('/api/base/captcha')
}