import { Link } from "react-router-dom"

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div>
      <header className="container nav">
        <div className="nav__content">
          <p><Link to="/">coffee</Link></p>
        </div>
        <div className="nav__content">
          <ul className="nav__content__links">
            <li><Link to="/setup">setup</Link></li>
            <li><Link to="/login">login</Link></li>
          </ul>
        </div>
      </header>
      <main>
        {children}
      </main>
    </div>
  )
}

export default Layout