'use client';

import clsx from 'clsx';
import {
    MouseEventHandler,
    useCallback,
    useEffect,
    useRef,
    useState,
    forwardRef,
    useImperativeHandle,
} from 'react';

type Props = {
    defaultWidth: number;
    minWidth?: number;
    maxWidth?: number;
    resizeDir: 'Left' | 'Right';
    onSizeChanged?: (w: number) => void;
} & React.DetailedHTMLProps<
    React.HTMLAttributes<HTMLDivElement>,
    HTMLDivElement
>;

export const ResizeContainer = forwardRef<HTMLDivElement, Props>(
    function ResizeContainerLeft(
        {
            className = '',
            children,
            defaultWidth,
            minWidth = 0,
            maxWidth = 9999,
            resizeDir,
            onSizeChanged,
            ...props
        }: Props,
        ref
    ) {
        const containerRef = useRef<HTMLDivElement>(null);

        const [width, setWidth] = useState(defaultWidth);
        const [offset, setOffset] = useState(0);
        const [isMouseDown, setIsMouseDown] = useState(false);

        useImperativeHandle(ref, () => containerRef.current!, []);

        const handleMouseDown: MouseEventHandler = useCallback((e) => {
            if (e.button === 0) {
                window?.getSelection()?.removeAllRanges();
                setIsMouseDown(true);
            }
            setIsMouseDown(true);
        }, []);

        useEffect(() => {
            onSizeChanged && onSizeChanged(width);
        }, [onSizeChanged, width]);

        // set width when min and max size
        useEffect(() => {
            setWidth((priv) =>
                priv > maxWidth ? maxWidth : priv < minWidth ? minWidth : priv
            );
        }, [maxWidth, minWidth]);

        useEffect(() => {
            if (!containerRef.current) return;
            if (!isMouseDown) return;

            const handleMouseMove = (e: MouseEvent) => {
                window?.getSelection()?.removeAllRanges();

                const rect = containerRef.current?.getBoundingClientRect();
                if (rect === undefined) {
                    return;
                }

                const offset =
                    resizeDir === 'Left'
                        ? e.x - rect.left - 2
                        : -(e.x - (rect.left + rect.width) + 3);
                const maxOffset = rect.width - minWidth;
                const minOffset = rect.width - maxWidth;
                setOffset(
                    offset < minOffset
                        ? minOffset
                        : offset > maxOffset
                        ? maxOffset
                        : offset
                );
            };

            document.addEventListener('mousemove', handleMouseMove);

            return () => {
                document.removeEventListener('mousemove', handleMouseMove);
            };
        }, [isMouseDown, maxWidth, minWidth, resizeDir]);

        useEffect(() => {
            if (!isMouseDown) return;

            const handleMouseClick = () => {
                setIsMouseDown(false);

                const rect = containerRef.current?.getBoundingClientRect();
                if (rect === undefined) {
                    return;
                }

                setWidth(rect.width - offset);
                setOffset(0);
            };

            document.addEventListener('mouseup', handleMouseClick);

            return () => {
                document.removeEventListener('mouseup', handleMouseClick);
            };
        }, [isMouseDown, offset]);

        return (
            <div
                ref={containerRef}
                className={`relative h-full ${className} ${
                    !isMouseDown && 'transition-all duration-300'
                }`}
                style={{ width: `${width}px` }}
                {...props}
            >
                <div
                    onMouseDown={handleMouseDown}
                    className={`absolute w-[4px] h-full top-0 cursor-w-resize hover:bg-primary z-[900] ${clsx(
                        {
                            'transition-all duration-300': !isMouseDown,
                            'bg-primary before:absolute before:w-4 before:h-full before:top-0 before:-left-2 before:cursor-w-resize before:opacity-0':
                                isMouseDown,
                        }
                    )}`}
                    style={
                        resizeDir === 'Left'
                            ? { left: `${offset}px` }
                            : { right: `${offset}px` }
                    }
                ></div>
                {children}
            </div>
        );
    }
);
