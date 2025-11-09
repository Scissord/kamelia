// import { Footer, SideBar, Header } from "@/components";

export default function ShopLayout({
  children,
}: {
  children: React.ReactNode;
}) {

  return (
    <>
      {/* <Header />
      <SideBar /> */}
      <main className="min-h-screen">
        {children}
      </main>
      {/* <Footer /> */}
    </>
  );
}
