import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Todo from './components/Todo';
import Auth from './components/Auth';
import axios from 'axios';
import { useEffect } from 'react';
import { CsrfToken } from './types';

const App = () => {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken().catch(e=>console.error(e))
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
