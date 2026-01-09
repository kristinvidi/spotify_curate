import { NextResponse } from 'next/server';
import { grpcAsync } from '@/lib/grpc';

export async function POST(request) {
    const body = await request.json();
    const { tag } = body;

    if (!tag) {
        return NextResponse.json({ error: 'Tag is required' }, { status: 400 });
    }

    try {
        // Backend expects a list of genres
        const response = await grpcAsync('CreatePlaylistRecentInGenre', { genre: [tag] });

        if (!response.general?.success) {
            return NextResponse.json(
                { error: response.general?.failure_details || 'Failed to create playlist' },
                { status: 500 }
            );
        }

        return NextResponse.json({ success: true });
    } catch (error) {
        console.error('gRPC Error:', error);
        return NextResponse.json({ error: error.message }, { status: 500 });
    }
}
