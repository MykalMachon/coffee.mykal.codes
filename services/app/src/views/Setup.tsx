import Layout from "../components/Layout";

const SetupPage = () => {
  return (
    <Layout>
      <h1>Setup</h1>
      <p>You haven't set the site up yet!</p>

      <p>Create your account below and you should be good to go</p>

      <form method="POST" action="/api/signup">
        <label>
          Username:
          <input type="text" name="username" />
        </label>
        <label>
          Password:
          <input type="password" name="password" />
        </label>
        <label>
          Confirm Password:
          <input type="password" name="confirmPassword" />
        </label>
        <button type="submit">Sign up</button>
      </form>
    </Layout>
  );
}

export default SetupPage;