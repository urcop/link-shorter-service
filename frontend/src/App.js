import Header from './components/Header'
import LinkInput from './components/LinkInput'
import style from './App.module.scss'

const App = () => {
  return (
      <div className={style.container}>
        <Header/>
        <LinkInput/>
      </div>
  );
}

export default App;
