import React, {useState} from 'react'
import {Form, Input, Button, Modal} from 'antd'
import {GithubOutlined} from '@ant-design/icons'
import {useLocation} from 'react-router-dom'

import {GITHUB} from '@/config'
import {save} from '@/utils/localStorage'

// redux
import {login, register} from '@/redux/user/actions'
import {useDispatch} from 'react-redux'

// hooks
import {useListener} from '@/hooks/useBus'
import {GetCaptcha} from "@/api/user";

const FormItemLayout = {
    labelCol: {
        xs: {span: 24},
        sm: {span: 6}
    },
    wrapperCol: {
        xs: {span: 24},
        sm: {span: 18}
    }
}

function FormItem(props) {
    const {children, ...rest} = props
    return <Form.Item {...FormItemLayout} {...rest}>{children}</Form.Item>
}

function SignModal(props) {
    const dispatch = useDispatch() // dispatch hooks
    const location = useLocation() // location
    const [captcha, setCaptcha] = useState({
        picPath: '',
        captchaId: ''
    })
    const [visible, setVisible] = useState(false)
    const [type, setType] = useState('login')
    const [form] = Form.useForm();
    const initCaptcha = () => {
        GetCaptcha().then(res => {
            if (res.code === 0) {
                setCaptcha(res.data)
            }
        })
    }
    useListener('openSignModal', async type => {
        form.resetFields()
        setType(type)
        if (type === 'login') {
            initCaptcha()
        }
        setVisible(true)
    })

    function handleSubmit(e) {
        e.preventDefault()
        form.validateFields().then(( values) => {
            const action = type === 'login' ? login : register
            const params = type === 'login' ? {...values, captchaId: captcha.captchaId} : values
            dispatch(action(params)).then((v) => {
                console.log(v)
                setVisible(false) // type =  login | register
            })
        }).catch(err => {
            console.error(err)
        })
    }

    function githubLogin() {
        const {pathname, search} = location
        save('prevRouter', `${pathname}${search}`)
        window.location.href = `${GITHUB.url}?client_id=${GITHUB.client_id}`
    }

    // 确认密码
    function compareToFirstPassword(rule, value, callback) {
        if (value && value !== form.getFieldValue('password')) {
            callback('Two passwords that you enter is inconsistent!')
        } else {
            callback()
        }
    }

    function handleRefreshCaptcha() {
       initCaptcha()
    }

    return (
        <Modal
            width={460}
            title={type}
            visible={visible}
            onCancel={e => setVisible(false)}
            footer={null}>
            <Form layout='horizontal' form={form}>
                {type === 'login' ? (
                        <>
                            <FormItem label='用户名' name="username"
                                      rules={[{required: true, message: 'Username is required'}]}>
                                <Input placeholder='请输入用户名'/>
                            </FormItem>
                            <FormItem label='密码' name="password"
                                      rules={[{required: true, message: 'Password is required'}]}>
                                <Input placeholder='请输入密码' type='password'/>
                            </FormItem>
                            <FormItem label='验证码' name="captcha"
                                      rules={[{required: true, message: 'captcha is required'}]}>
                                <Input placeholder='请输入验证码' type='text'/>

                            </FormItem>
                            <FormItem label={' '}>
                                <img src={captcha.picPath} onClick={handleRefreshCaptcha} alt={captcha.captchaId}/>
                            </FormItem>
                        </>
                    )
                    : (
                        <>
                            <FormItem label='用户名' name={'username'}
                                      rules={[{required: true, message: 'Username is required'}]}>
                                <Input placeholder='请输入用户名'/>
                            </FormItem>
                            <FormItem label='密码' name={'password'} rules={[{required: true, message: 'Password is required'}]}>
                                <Input placeholder='请输入密码' type='password'/>
                            </FormItem>
                            <FormItem label='确认密码' name={'confirm'} rules={[
                                {required: true, message: 'Password is required'},
                                {validator: compareToFirstPassword}
                            ]}>
                                <Input placeholder='确认密码' type='password'/>
                            </FormItem>
                            <FormItem label='邮箱' name="email" rules={[
                                {type: 'email', message: 'The input is not valid E-mail!'},
                                {required: true, message: 'Please input your E-mail!'}
                            ]}>
                                <Input placeholder='请输入您的邮箱'/>
                            </FormItem>
                        </>
                    )}
            </Form>
            <Button type='primary' block onClick={handleSubmit}>
                {type}
            </Button>
            {GITHUB.enable && (
                <Button block icon={<GithubOutlined/>} onClick={githubLogin} style={{marginTop: 10}}>
                    github login
                </Button>
            )}
        </Modal>
    )
}


export default SignModal
