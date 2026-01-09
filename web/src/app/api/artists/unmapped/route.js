import { NextResponse } from 'next/server';
import { grpcAsync } from '@/lib/grpc';

export async function GET(request) {
    const { searchParams } = new URL(request.url);
    const userID = searchParams.get('userID');

    if (!userID) {
        return NextResponse.json({ error: 'UserID is required' }, { status: 400 });
    }

    try {
        const response = await grpcAsync('GetUnmappedArtistsForUser', { user_spotify_id: userID });

        if (!response.general?.success) {
            return NextResponse.json(
                { error: response.general?.failure_details || 'Failed to fetch artists' },
                { status: 500 }
            );
        }

        return NextResponse.json({ artists: response.artists || [] });
    } catch (error) {
        console.error('gRPC Error:', error);
        return NextResponse.json({ error: error.message }, { status: 500 });
    }
}
