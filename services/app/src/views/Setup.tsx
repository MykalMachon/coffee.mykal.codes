import Layout from "../components/Layout";
import { API_URL } from "../utils/api";

const SetupPage = () => {

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
      <h1>Setup</h1>
      <p>You haven't set the site up yet!</p>

      <p>Create your account below and you should be good to go</p>

      <form method="POST" action={`${API_URL}/auth/signup`} onSubmit={handleSubmit}>
        <label>
          Full Name:
          <input type="text" name="name" />
        </label>
        <label>
          Email:
          <input type="text" name="email" />
        </label>
        <label>
          Password:
          <input type="password" name="password" />
        </label>
        <label>
          Confirm Password:
          <input type="password" name="passwordConfirmation" />
        </label>
        <button type="submit">Sign up</button>
      </form>
    </Layout>
  );
}

export default SetupPage;