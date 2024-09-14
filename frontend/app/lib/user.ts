export interface UserDetails {
  id: number;
  email: string;
  created_at: Date;
  updated_at: Date;
}

export async function getUserDetails(): Promise<UserDetails | null> {
  const token = localStorage.getItem('token');
  if (!token) {
    return null;
  }
  const response = await fetch(process.env.API_HOST + '/api/v1/users', {
    headers: {
      'Authorization': 'Bearer ' + token!,
    },
  });
  if (!response.ok) {
    return null;
  }
  const data = await response.json();
  return data;
}