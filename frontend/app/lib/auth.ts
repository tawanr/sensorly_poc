export interface TokenData {
  token: string;
  expiry: Date;
};

export async function loginUser(email: string, password: string): Promise<string> {
  const response = await fetch("http://localhost:8000" + '/api/v1/tokens/authentication', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email,
      password,
    }),
  });
  if (!response.ok) {
    throw new Error('Failed to login');
  }
  const data: TokenData = await response.json();
  return data.token;
}
