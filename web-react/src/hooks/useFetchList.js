import {useEffect, useState, useCallback} from 'react'

import {useLocation, useHistory} from 'react-router-dom'
import {decodeQuery} from '@/utils'
import useMount from './useMount'

/**
 * fetchList
 * requestUrl 请求地址
 * queryParams 请求参数
 * withLoading 是否携带 loading
 * fetchDependence 依赖 => 可以根据地址栏解析拉取列表
 */
export default function useFetchList(
    {
        requestMethod = new Promise(),
        queryParams = null,
        usePathParams = true,
        withLoading = true,
        fetchDependence = []
    }
) {
    const [list, setDataList] = useState([])
    const [loading, setLoading] = useState(false)
    const [pagination, setPagination] = useState({pageNo: 1, pageSize: 10, total: 0})

    const location = useLocation()
    const history = useHistory()

    useMount(() => {
        if (fetchDependence.length === 0) {
            fetchWithLoading()
        }
    })

    useEffect(() => {
        if (fetchDependence.length > 0) {
            const params = usePathParams ? decodeQuery(location.search) : {}
            fetchWithLoading(params)
        }
    }, fetchDependence) // eslint-disable-line  react-hooks/exhaustive-deps

    function fetchWithLoading(params) {
        withLoading && setLoading(true)
        fetchDataList(params)
    }

    function fetchDataList(params) {
        const requestParams = {
            pageNo: pagination.pageNo,
            pageSize: pagination.pageSize,
            ...queryParams,
            ...params
        }
        requestParams.pageNo = parseInt(requestParams.pageNo)
        requestParams.pageSize = parseInt(requestParams.pageSize)
        requestParams.category = parseInt(requestParams.category)
        requestParams.orderType = parseInt(requestParams.orderType)
        requestMethod && requestMethod(requestParams).then(res => {
            if (res.code === 0) {
                let data = res.data
                pagination.pageNo = parseInt(requestParams.pageNo)
                pagination.pageSize = parseInt(requestParams.pageSize)
                pagination.category = parseInt(requestParams.category)
                pagination.orderType = parseInt(requestParams.orderType)
                pagination.total = data.total
                setPagination({...pagination})
                setDataList(data.list || [])
            }
            withLoading && setLoading(false)
        }).catch(e => withLoading && setLoading(false))
    }

    const onFetch = useCallback(
        params => {
            withLoading && setLoading(true)
            fetchDataList(params)
        },
        // eslint-disable-next-line react-hooks/exhaustive-deps
        [queryParams]
    )

    const handlePageChangeByPath = useCallback(
        pageNo => {
            // return
            const search = location.search.includes('pageNo=')
                ? location.search.replace(/(pageNo=)(\d+)/, `$1${pageNo}`)
                : `?pageNo=${pageNo}`
            const jumpUrl = location.pathname + search
            history.push(jumpUrl)
        },
        // eslint-disable-next-line react-hooks/exhaustive-deps
        [queryParams, location.pathname]
    )
    const handlePageChange = pageNo => {
        fetchWithLoading({pageNo})
    }

    return {
        list,
        loading,
        pagination: {
            ...pagination,
            onChange: usePathParams ? handlePageChangeByPath : handlePageChange
        },
        onFetch
    }
}
