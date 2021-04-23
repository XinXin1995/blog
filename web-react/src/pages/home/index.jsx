import ArticleList from '@/pages/articleList/withContent'
import './index.less'
import '@/assets/style/article.less'

function Home () {
  return (
    <div className={'home'}>
        <ArticleList />
    </div>
  )
}

export default Home
