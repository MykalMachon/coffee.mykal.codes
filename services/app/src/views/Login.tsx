import Layout from "../components/Layout";
import { API_URL } from "../utils/api";

const LoginPage = () => {

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;
    const response = await fetch(form.action, {
      method: form.method,
      body: new FormData(form),
    });
    if (response.ok){ console.log('account created!') }
    else { console.error('account creation failed!') }
  }

  return (
    <Layout>
      <h1>Login</h1>
      <p>You need to login to write posts. If you're not Mykal good luck!</p>
      <form method="POST" action={`${API_URL}/auth/login`} onSubmit={handleSubmit}>
        <label>
          Username:
          <input type="text" name="email" />
        </label>
        <label>
          Password:
          <input type="password" name="password" />
        </label>
        <button type="submit">Login</button>
      </form>
    </Layout>
  );
}

export default LoginPage;