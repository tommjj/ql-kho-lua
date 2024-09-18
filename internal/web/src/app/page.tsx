import NavItem from '@/components/nav/nav-items';

export default function Home() {
    return (
        <div className="h-screen w-screen ">
            <NavItem active href="/map"></NavItem>
        </div>
    );
}
