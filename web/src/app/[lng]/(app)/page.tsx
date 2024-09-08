export default function Home({params: {lng}}: {params: {lng: string}}) {
  return (
    <>
      Home {lng}
        <button className="bg-primary-500">click</button>
    </>
  );
}
