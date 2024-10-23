import { Button } from '@/components/shadcn-ui/button';
import { ArrowRight, Box } from 'lucide-react';
import Link from 'next/link';

export default function Home() {
    return (
        <section>
            <header className="absolute top-0 left-0 right-0 flex items-center justify-between px-4 border-b">
                <div className="flex items-center text-2xl font-semibold text-opacity-90 my-2 h-9">
                    <Box className="size-8 mr-2 text-primary"></Box>
                    Rice Warehouses
                </div>
                <div>
                    <Button className="px-11" asChild>
                        <Link href={'/log-in'}>Login</Link>
                    </Button>
                </div>
            </header>

            <section className="snap-always snap-center relative z-20 flex w-full h-screen overflow-hidden">
                <div className="flex relative z-10 items-start flex-col lg:text-left justify-center w-full lg:w-full h-full">
                    <div className="relative pb-12  pt-32 lg:pt-0 flex flex-col justify-center h-full lg:h-auto lg:block w-full p-4 lg:pb-4 md:px-12 lg:px-24">
                        <div></div>
                        <div>
                            <h1 className="text-5xl tracking-tighter lg:mt-0 text-[#333] dark:text-white font-semibold lg:top-0 md:text-6xl md:font-semibold lg:text-[4.8rem] lg:-ml-1 leading-none">
                                Organize your work
                            </h1>

                            <p className="tracking-tight max-w-[340px] sm:max-w-none text-xl sm:text-2xl lg:px-0 mt-7 lg:mt-3 text-[#666] dark:text-white ">
                                A rice warehouse manager
                            </p>
                        </div>
                        <div></div>
                        <div className="flex justify-center mt-7 mb-5 lg:mb-0 lg:justify-start relative w-full lg:mt-10 sm:max-w-[640px]">
                            <div>
                                <Button
                                    asChild
                                    className="mt-4 sm:mt-0 h-12 w-full flex-shrink flex-grow-0 group sm:w-44 px-7 pl-8 py-2 "
                                >
                                    <Link href={'/log-in'}>
                                        Get Started
                                        <ArrowRight className="h-6 w-6 text-white group-hover:translate-x-1 transition-transform" />
                                    </Link>
                                </Button>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </section>
    );
}
