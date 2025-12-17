import { Header } from './header';
import TableWrapper from './table_wrapper';

export default async function Home() {
  const res = await fetch('http://localhost:8080/api/clients', {
    cache: 'no-store', // без кеша
  });

  const data = await res.json();

  return (
    <div className="p-2">
      <Header />
      <TableWrapper data={data} />
    </div>
  );
}
